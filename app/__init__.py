# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.
from __future__ import division
from PySide.QtGui import QApplication, QLabel, QWidget, QMessageBox
from PySide.QtGui import QStatusBar, QProgressBar, QPushButton
from PySide.QtGui import QVBoxLayout, QHBoxLayout, QFileDialog
from PySide.QtGui import QDesktopWidget, QInputDialog, QLineEdit
from PySide.QtGui import QIcon, QFont, QPixmap
from PySide.QtCore import Qt, Slot, QEvent, QUrl
from sys import exit, platform, argv
from os import name, system, path, remove
from sys import platform as sysname
from sqlalchemy import create_engine, ForeignKey
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, Boolean, Binary
from sqlalchemy.orm import sessionmaker
from threads import EncryptTH, DecryptTH
from ex_functions import encryptit, isenct, r_path
from hashlib import sha256


class SafeLock(QWidget):
    def __init__(self):
        super(SafeLock, self).__init__()
        main_layout = QVBoxLayout(self)
        self.Version = '0.5 beta'
        self.s_error = "QStatusBar{color:red;font-weight:1000;}"
        self.s_loop = "QStatusBar{color:black;font-weight:1000;}"
        self.s_norm = "QStatusBar{color:blue;font-style:italic;"
        self.s_norm += "font-weight:500;}"
        self.Processing = None
        self.CP = None
        self.PPbar = None
        self.key = None
        self.DFiles = []

        self.picon = r_path("images/favicon.png")
        self.plogo = r_path("images/logo.png")
        if name == 'nt':
            self.picon = r_path("images\\favicon.png")
            self.plogo = r_path("images\\logo.png")
        self.icon = QIcon(self.picon)
        self.logo = QIcon(self.plogo)
        self.center()
        self.setStyle()
        self.setMW(main_layout)
        self.setSB(main_layout)
        self.show()

    def db(self, password="password", dbname="untitled.sld", dc=True):
        eng = create_engine("sqlite:///%s" % dbname,
                            connect_args={'check_same_thread': False})
        Base = declarative_base(bind=eng)
        Session = sessionmaker(bind=eng)
        session = Session()

        class Identifier(Base):
            __tablename__ = "identifier"
            id = Column(Integer, primary_key=True)
            version = Column(Binary)
            kvd = Column(Binary)

            def __init__(self, version, kvd):
                self.id = 0
                self.version = version
                self.kvd = kvd

        class Folder(Base):
            __tablename__ = 'folders'
            id = Column(Integer, primary_key=True)
            path = Column(String)

            def __init__(self, path="Empty"):
                self.path = path

        class File(Base):
            __tablename__ = 'files'
            id = Column(Integer, primary_key=True)
            name = Column(String)
            f_id = Column(Integer, ForeignKey('folders.id'), nullable=True)
            bb = Column(Binary)

            def __init__(self, name="empty", f_id=0, bb=1010):
                self.name = name
                self.f_id = f_id
                self.bb = bb
        if dc:
            Base.metadata.create_all()
            enc = encryptit(sha256(self.Version).digest(),
                            sha256(password).digest())
            session.add(Identifier(enc[0], enc[1]))
            session.commit()

        return [eng, Base, session, File, Folder, Identifier, password]

    def checkP(self, db):
        try:
            d = db[2].query(db[5]).filter_by(id=0).first()
            if d is not None:
                if isenct(self.Version, d.version, db[6], d.kvd):
                    return True
        except:
            pass
        return False

    def setStyle(self):
        self.setMaximumWidth(410)
        self.setMinimumWidth(410)
        self.setMaximumHeight(370)
        self.setMinimumHeight(370)
        self.setWindowIcon(self.icon)
        self.activateWindow()
        self.setWindowTitle("safelock " + self.Version)
        self.setToolTip(
            "Just drag and drop any files or folders" +
            " to ecrypt or a .sld file to decrypt")
        self.setAcceptDrops(True)
        self.show()

    def center(self):
        qrect = self.frameGeometry()
        cp = QDesktopWidget().availableGeometry().center()
        qrect.moveCenter(cp)
        self.move(qrect.topLeft())

    def setMW(self, ml):
        ll = QVBoxLayout()
        self.mIcon = QLabel()
        self.mIcon.setAlignment(Qt.AlignCenter)
        mmicon = self.logo.pixmap(250, 230, QIcon.Active, QIcon.On)
        self.mIcon.setPixmap(mmicon)
        self.mInst = QLabel(
            "<center>(Drag and drop files or folders to encrypt them)<br>" +
            "(Drap and drop .sld file to decrypt it)<br>" +
            "<u>2GB max single file size to encrypt</u></center>")
        font = QFont()
        font.setPointSize(13)
        font.setBold(True)
        font.setWeight(75)
        self.fInst = QLabel('<center>| Double-Click for about |</center>')
        self.mInst.setFont(font)
        ll.addWidget(self.mIcon)
        ll.addWidget(self.mInst)
        ll.addWidget(self.fInst)
        ml.addLayout(ll)

    def setSB(self, ml):
        self.statusb = QStatusBar()
        ml.addWidget(self.statusb)

    def modSB(self):
        if self.PPbar is None:
            self.PPbar = True
            self.statusb.clear()
            self.pbar = QProgressBar()
            self.pbar.setMaximum(100)
            self.pbar.setMinimum(0)
            self.plabel = QLabel()
            self.statusb.addWidget(self.plabel, 0.5)
            self.statusb.addWidget(self.pbar, 3)
        else:
            self.PPbar = None
            self.statusb.removeWidget(self.plabel)
            self.statusb.removeWidget(self.pbar)
            self.statusb.clear()

    def saveFile(self, fl):
        fname, _ = QFileDialog.getSaveFileName(self,
                                               "Save encrypted file",
                                               fl,
                                               "Safelock file (*.sld)")
        if '.' in fname:
            tm = fname.split('.')
            tm = tm[len(tm) - 1]
            if tm == "sld":
                try:
                    if path.isfile(fname):
                        remove(fname)
                except:
                    pass
                return fname
        if len(fname) <= 0:
            return None
        fname += ".sld"
        try:
            if path.isfile(fname):
                remove(fname)
        except:
            pass
        return fname

    def extTo(self, fl):
        fname = QFileDialog.getExistingDirectory(self, "Extract files to",
                                                 fl)
        if len(fname) <= 0:
            return None
        return fname

    def getPass(self):
        passwd, re = QInputDialog.getText(self, "Password", "Enter password :",
                                          QLineEdit.Password)
        if not re:
            return False
        if len(passwd) <= 0:
            return None
        return passwd

    def errorMsg(self, msg="Something went wrong !"):
        QMessageBox.critical(self, "Error", msg)
        return True

    def aboutMsgg(self):
        Amsg = "<center>All credit reserved to the author of "
        Amsg += "safelock %s " % self.Version
        Amsg += ", This work is a free, open-source project licensed "
        Amsg += " under Mozilla Public License version 2.0 . <br><br>"
        Amsg += " visit for more info or report:<br> "
        Amsg += "<b><a href='https://safe-lock.github.io'> "
        Amsg += "https://safe-lock.github.io/ </a> </b></center>"
        QMessageBox.about(self,
                          "About",
                          Amsg
                          )
        return True

    def getSession(self):
        self.eng = eng(self.password, self.dbname)
        Session = sessionmaker(bind=self.eng)
        self.Base = declarative_base(bind=self.eng)
        self.Base.metadata.create_all()
        self.session = Session()
        return self.session

    def dragEnterEvent(self, e):
        if e.mimeData().hasUrls:
            e.accept()
        else:
            e.ignore()

    def dragMoveEvent(self, e):
        if e.mimeData().hasUrls:
            e.accept()
        else:
            e.ignore()

    def dropEvent(self, e):
        if e.mimeData().hasUrls:
            e.setDropAction(Qt.CopyAction)
            e.accept()
            self.DFiles = []
            for url in e.mimeData().urls():
                try:
                    if sysname == 'darwin':
                        from Foundation import NSURL
                        fname = NSURL.URLWithString_(
                            url.toString()).filePathURL().path()
                    else:
                        fname = url.toLocalFile()
                    self.DFiles.append(fname)
                except:
                    pass
            self.dealD(self.DFiles)
        else:
            event.ignore()

    def in_loop(self):
        if self.Processing is None:
            self.Processing = True
            self.setAcceptDrops(False)
            self.mIcon.setEnabled(False)
            self.mInst.setEnabled(False)
            self.fInst.setText("<center>| Double-Click to cancel |</center>")
            self.setToolTip("Double-Click anywhere to cancel")
        else:
            self.Processing = None
            self.setAcceptDrops(True)
            self.mIcon.setEnabled(True)
            self.mInst.setEnabled(True)
            self.fInst.setText("<center>| Double-Click for about |</center>")
            self.setToolTip(
                "Just drag and drop any files or folders" +
                " to ecrypt or a .sld file to decrypt")

    def dealD(self, files):
        def tpp(inp):
            a = path.basename(inp)
            return inp.replace(a, '')
        if len(files) < 1:
            return False
        elif len(files) >= 1:
            if len(files) == 1:
                tf = files[0].split('.')
                tf = tf[len(tf) - 1]
                if tf == 'sld':
                    pw = self.getPass()
                    if pw is None:
                        self.errorMsg("You can't set an empty password !")
                        return False
                    elif not pw:
                        return False
                    if not self.checkP(self.db(sha256(pw).digest(), files[0],
                                               dc=False)):
                        self.errorMsg(
                            "Wrong password entered, try again.")
                        return False
                    else:
                        fold = self.extTo(tpp(files[0]))
                        if fold is not None:
                            self.CP = fold
                            self.in_loop()
                            self.P = DecryptTH(fold, self.db(pw,
                                                             files[0],
                                                             dc=False),
                                               sha256(pw).digest())
                            self.P.start()
                            self.P.somesignal.connect(self.handleStatusMessage)
                            self.P.setTerminationEnabled(True)
                            return True
            pw = self.getPass()
            if pw is None:
                self.errorMsg("You can't set an empty password !")
            elif not pw:
                pass
            else:
                fil = self.saveFile(tpp(files[0]))
                if fil is not None:
                    if path.isfile(fil):
                        try:
                            remove(fil)
                        except:
                            pass
                    self.CP = fil
                    self.in_loop()
                    self.P = EncryptTH(files, self.db(pw, fil),
                                       sha256(pw).digest())
                    self.P.start()
                    self.P.somesignal.connect(self.handleStatusMessage)
                    self.P.setTerminationEnabled(True)
        return True

    @Slot(object)
    def handleStatusMessage(self, message):
        self.statusb.setStyleSheet(self.s_loop)

        def getIT(f, o):
            return int((f / o) * 100)
        mm = message.split('/')
        if mm[len(mm) - 1] == '%':
            if self.PPbar is None:
                self.modSB()
            self.pbar.setValue(getIT(int(mm[0]), int(mm[1])))
            self.setWindowTitle("Processing : " + str(getIT(int(mm[0]),
                                                            int(mm[1]))) + "%")
            self.plabel.setText(mm[0] + '/' + mm[1])
        else:
            self.unsetCursor()
            if self.PPbar is not None:
                self.modSB()
            if message[:7] == '# Error':
                if self.Processing:
                    self.in_loop()
                self.statusb.setStyleSheet(self.s_error)
                self.cleanup()
            elif message[:6] == '# Stop':
                self.statusb.setStyleSheet(self.s_error)
            elif message[:6] == '# Done':
                if self.Processing:
                    self.in_loop()
                self.statusb.setStyleSheet(self.s_norm)
            elif message == "# Loading":
                self.setCursor(Qt.BusyCursor)
                self.statusb.setStyleSheet(self.s_norm)
            self.setWindowTitle('safelock ' + self.Version)
            self.statusb.showMessage(message)

    def mousePressEvent(self, event):
        if event.type() == QEvent.Type.MouseButtonDblClick:
            if self.Processing is None:
                self.aboutMsgg()
            else:
                self.closeEvent()

    def closeEvent(self, event=None):
        if self.Processing is not None:
            if event is not None:
                r = QMessageBox.question(
                    self,
                    "Making sure",
                    "Are you sure you want to exit, during an active" +
                    " process ?",
                    QMessageBox.Yes | QMessageBox.No)
            else:
                r = QMessageBox.question(
                    self,
                    "Making sure",
                    "Are you sure, you to cancel ?",
                    QMessageBox.Yes | QMessageBox.No)
            if r == QMessageBox.Yes:
                self.P.stop()
                self.cleanup()
                if event is not None:
                    self.in_loop()
                    event.accept()
                else:
                    self.in_loop()
            else:
                if event is not None:
                    event.ignore()
        else:
            if self.CP is not None:
                try:
                    if path.isfile(self.CP + '-journal'):
                        remove(self.CP + '-journal')
                except:
                    pass
            if event is not None:
                event.accept()

    def cleanup(self):
        if self.CP is not None:
            try:
                if path.isfile(self.CP + '-journal'):
                    remove(self.CP + '-journal')
                if path.isfile(self.CP):
                        remove(self.CP)
            except:
                pass


def gui():
    app = QApplication(argv)
    window = SafeLock()
    window.show()
    app.exec_()
