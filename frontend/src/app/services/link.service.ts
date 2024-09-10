import { Injectable } from '@angular/core';
import { defer, from, Observable } from 'rxjs';

import { GetVersion, ShowErrMsg } from '../../../wailsjs/go/backend/App';

export enum AppEvents {
  statusUpdateKey = 'status_update',
  statusEndKey = 'status_end',
  openedSlaFile = 'opened_sla_file',
}

@Injectable({
  providedIn: 'root'
})
export class LinkService {
  readonly appVersion$ = defer(() => from(GetVersion()));
  readonly files$ = new Observable<string[]>(sub => {
    window?.runtime?.OnFileDrop((x,y, fs) => sub.next(fs), false);
    window?.runtime?.EventsOn(AppEvents.openedSlaFile, f => sub.next([f]));
  });

  openUrl(url: string) {
    window?.runtime?.BrowserOpenURL(url);
  }

  removeDroppedListeners() {
    window?.runtime?.OnFileDropOff();
  }

  showError(msg: string) {
    ShowErrMsg(msg);
  }
}
