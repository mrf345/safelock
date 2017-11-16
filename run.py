# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.
from app import gui
from sys import exit, exc_info

if __name__ == '__main__':
    try:
        gui()
    except Exception:
        print(exc_info()[1])
        print("Error: just crached, please help us improve by reporting to us")
        print("\n\thttps://github.com/mrf345/safelock")
    exit(0)
