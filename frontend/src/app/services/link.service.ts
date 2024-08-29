import { Injectable } from '@angular/core';
import { defer, from, Observable } from 'rxjs';

import { GetVersion, ShowErrMsg } from '../../../wailsjs/go/main/App';

@Injectable({
  providedIn: 'root'
})
export class LinkService {
  readonly appVersion$ = defer(() => from(GetVersion()));
  readonly files$ = new Observable<string[]>(sub =>
    window?.runtime?.OnFileDrop((x,y, fs) => sub.next(fs), false)
  );

  openUrl(url: string) {
    window?.runtime?.BrowserOpenURL(url);
  }

  setTitle(title: string) {
    window?.runtime?.WindowSetTitle(title);
  }

  removeDroppedListeners() {
    window?.runtime?.OnFileDropOff();
  }

  showError(msg: string) {
    ShowErrMsg(msg);
  }
}
