from PySide.QtCore import QThread, Signal
from os import walk, path, remove
from pathlib import Path
from time import sleep
from random import randint
from struct import pack, unpack, calcsize
from Crypto.Cipher import AES
from tempfile import NamedTemporaryFile
from ex_functions import writeit


class EncryptTH(QThread):
    somesignal = Signal(object)

    def __init__(self, inp=[], db=[], key=None):
        QThread.__init__(self)
        self.inp = inp
        self.session = db[2]
        self.File = db[3]
        self.Folder = db[4]
        self.key = key
        self.estm = 0
        for p in self.inp:
            if path.isdir(p):
                for sd, d, f in walk(p):
                    for fd in f:
                        if path.isfile(path.join(sd, fd)):
                            self.estm += 1
            else:
                self.estm += 1
            self.somesignal.emit("%% Loading files : " + str(self.estm))
        self.abo = None
        self.err = None

    def __del__(self):
        self.quit()
        self.terminate()

    def run(self):
        def torp(r, rp):
            return path.join(
                path.basename(r),
                rp).replace(
                    r.replace(path.basename(r), ''), '')

        def gsize(f):
            if int(path.getsize(f) / 100000) >= 2000:
                return False
            return True
        counter = 0
        for mainf in self.inp:
            if self.abo is not None:
                self.somesignal.emit("# Stop: got canceled")
                self.err = True
                break
            if self.err is not None:
                break
            if path.isdir(mainf):
                self.session.add(self.Folder(path.basename(mainf)))
                for sd, d, ff in walk(mainf):
                    if self.abo is not None:
                        self.somesignal.emit("# Stop: got canceled")
                        self.err = True
                        break
                    for fff in ff:
                        if not gsize(path.join(sd, fff)):
                            self.somesignal.emit(
                                '# Error: file is too large')
                            self.err = True
                            break
                        try:
                            counter += 1
                            cdd = self.session.query(
                                self.Folder).filter_by(
                                    path=torp(mainf, sd)).first()
                            if cdd is None:
                                self.session.add(self.Folder(torp(mainf, sd)))
                            cdd = self.session.query(
                                self.Folder).filter_by(
                                    path=torp(mainf, sd)).first()
                            self.session.add(self.File(
                                fff,
                                cdd.id,
                                writeit(path.join(sd, fff),
                                        path.join(mainf, fff + '.tmp'),
                                        self.key)))
                            self.session.commit()
                            remove(path.join(mainf, fff + '.tmp'))
                            self.somesignal.emit(str(counter) + '/' +
                                                 str(self.estm) + '/%')
                        except:
                            self.somesignal.emit(
                                '# Error: inaccessiblity or write issue')
                            self.err = True
                            break
            elif path.isfile(mainf):
                if not gsize(mainf):
                    self.somesignal.emit(
                        '# Error: file is too large')
                    self.err = True
                    break
                try:
                    counter += 1
                    self.session.add(self.File(
                        path.basename(mainf),
                        None,
                        writeit(mainf, mainf + '.tmp', self.key)))
                    self.session.commit()
                    remove(mainf + '.tmp')
                    self.somesignal.emit(str(counter) + '/' +
                                         str(self.estm) + '/%')
                except:
                    self.somesignal.emit(
                        '# Error: too large or inaccessible file')
                    self.err = True
                    break
            else:
                self.somesignal.emit(
                    "# Error: something bad went wrong")
                self.err = True
                break
        if self.err is None:
            self.somesignal.emit("# Done: all encrypted")

    def stop(self):
        self.abo = True
        return True


class DecryptTH(QThread):
    somesignal = Signal(object)

    def __init__(self, path=None, db=[], key=None):
        QThread.__init__(self)
        self.path = path
        self.session = db[2]
        self.File = db[3]
        self.Folder = db[4]
        self.key = key
        self.estmf = self.session.query(self.Folder).count()
        self.estm = self.session.query(self.File).count()
        self.abo = None
        self.err = None

    def __del__(self):
        self.quit()
        self.terminate()

    def run(self):
        counter = 0
        for f in self.session.query(self.Folder):
            if self.abo is not None:
                self.somesignal.emit("# Stop: got canceled")
                self.err = True
                break
            try:
                counter += 1
                if not path.isdir(path.join(self.path, f.path)):
                    Path(path.join(self.path, f.path)).mkdir(parents=True)
                    self.somesignal.emit(
                        str(counter) + '/' + str(self.estmf) + '/%')
                    sleep(0.01)
            except:
                self.somesignal.emit("# Error: permissions or overwrite issue")
                self.err = True
                break
        counter = 0
        self.somesignal.emit('# Loading')
        if self.err is not None:
            return False
        for f in self.session.query(self.File):
            if self.abo is not None:
                self.somesignal.emit("# Stop: got canceled")
                self.err = True
                break
            try:
                counter += 1
                if f.f_id is None:
                    ffp = path.join(self.path, f.name)
                else:
                    fol = self.session.query(
                        self.Folder).filter_by(id=f.f_id).first()
                    ffp = path.join(self.path, fol.path)
                    ffp = path.join(ffp, f.name)
                if path.isfile(ffp):
                    remove(ffp)
                tmpo = path.join(self.path, path.basename(self.path))
                with open(tmpo + '.tmp', 'wb') as ffpp:
                    ffpp.write(f.bb)
                    ffpp.close()
                with open(ffp, 'wb') as outfile:
                    with open(tmpo + '.tmp', 'rb') as inb:
                        chs = 24 * 1024
                        o_s = unpack('<Q', inb.read(calcsize('Q')))[0]
                        iv = inb.read(16)
                        decryptor = AES.new(self.key, AES.MODE_CBC, iv)
                        while True:
                            ch = inb.read(chs)
                            if len(ch) == 0:
                                break
                            outfile.write(decryptor.decrypt(ch))
                        outfile.truncate(o_s)
                        inb.close()
                    outfile.close()
                remove(tmpo + '.tmp')
                self.somesignal.emit(
                    str(counter) + '/' + str(self.estm) + '/%')
                sleep(0.1)
            except:
                self.somesignal.emit("# Error: permissions or overwrite issue")
                self.err = True
                break
        if self.err is None:
            self.somesignal.emit("# Done: all decrypted")

    def stop(self):
        self.abo = True
        return True
