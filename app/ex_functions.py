from random import randint
from struct import pack, unpack, calcsize
from Crypto.Cipher import AES
from os import path, remove
from hashlib import sha256
import sys


def writeit(ff, inf, key):
    if path.isfile(inf):
        try:
            remove(ff)
        except:
            pass
    with open(ff, 'rb') as mainfof:
        chs = 64 * 1024
        iv = ''.join(chr(randint(0, 0xFF)) for i in range(16))
        encryptor = AES.new(key, AES.MODE_CBC, iv)
        filesize = path.getsize(ff)
        with open(inf, 'wb') as otf:
            otf.write(pack('<Q', filesize))
            otf.write(iv)
            while True:
                ch = mainfof.read(chs)
                if len(ch) == 0:
                    break
                elif len(ch) % 16 != 0:
                    ch += ' ' * (16 - len(ch) % 16)
                otf.write(encryptor.encrypt(ch))
            otf.close()
        mainfof.close()
    with open(inf, 'rb') as otf:
        return otf.read()


def encryptit(text, key):
    if len(text) == 32 and len(key) == 32:
        iv = ''.join(chr(randint(0, 0xFF)) for i in range(16))
        enc = AES.new(key, AES.MODE_CBC, iv)
        return [enc.encrypt(text), iv]
    return []


def isenct(text, has, key, iv):
    for ar in [has, key, iv]:
        if len(ar) != 32 and len(ar) != 16:
            return False
    enc = AES.new(key, AES.MODE_CBC, iv)
    if enc.decrypt(has) == sha256(text).digest():
        return True
    return False


def r_path(relative_path):
    try:
        base_path = sys._MEIPASS
    except Exception:
        base_path = path.abspath(".")
    return path.join(base_path, relative_path)
