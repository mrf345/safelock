import { Injectable } from '@angular/core';
import { BehaviorSubject, from, Observable, tap } from 'rxjs';
import { TitleCasePipe } from '@angular/common';

import { LinkService, AppEvents } from './link.service';
import { Task, TaskKind } from '../helpers/task.class'
import { Encrypt, Decrypt, Cancel } from '../../../wailsjs/go/backend/App';


@Injectable({
  providedIn: 'root'
})
export class TaskService {
  private _task = new Task();
  private titleCase = new TitleCasePipe().transform

  done$ = new BehaviorSubject(false);
  task$ = new BehaviorSubject(this._task);

  get task() {
    return this.task$.getValue();
  }

  constructor(
    private readonly linkService: LinkService,
  ) {}

  create(files: string[])  {
    if (files.length > 1 && files.every(p => p.endsWith('.sla'))) {
      this.remove()
      this.linkService.showError(
        "You cannot decrypt multiple .sla files at once!"
      );
      return
    }

    this._task.files$.next(files);
    this._task.kind$.next(
      files.length === 1 && files[0].endsWith('.sla')
        ? TaskKind.decrypt
        : TaskKind.encrypt
    );
  }

  run(): Observable<string> {
    const password = this._task.password$.getValue();
    const kind = this._task.kind$.getValue();
    const files = this._task.files$.getValue();

    if (!password) {
      throw Error('Ran task without password')
    }

    if (kind === TaskKind.encrypt) {
      return from(Encrypt(files, password))
        .pipe(tap(id => {
          if (id) {
            this._task?.id$.next(id);
            this.startStatusUpdate();
          } else this.remove()
        }));
    } else if (kind === TaskKind.decrypt) {
      return from(Decrypt(files[0], password))
        .pipe(tap(id => {
          if (id) {
            this._task?.id$.next(id);
            this.startStatusUpdate();
          } else this.remove()
        }));
    }

    throw Error(`Unknown task ${kind}`);
  }

  cancel(): Observable<void> {
    this.restoreDefaultTitle();
    return from(Cancel()).pipe(tap(() => this.remove()));
  }

  remove() {
    this.stopStatusUpdate();
    this._task = new Task();
    this.task$.next(this._task);
    this.restoreDefaultTitle();
  }

  setPassword(pwd: string) {
    this._task.password$.next(pwd);
  }

  stopStatusUpdate() {
    window?.runtime?.EventsOff(AppEvents.statusUpdateKey);
    window?.runtime?.EventsOff(AppEvents.statusEndKey);
  }

  private startStatusUpdate() {
    window?.runtime?.EventsOn(AppEvents.statusUpdateKey, (status: string, percent: string) => {
      this._task?.status$.next(status);
      this._task?.percent$.next(percent);
      this.linkService.setTitle(
        `${this.titleCase(this._task?.kind$.getValue())}ing` +
        ` (${percent}%)`
      );
    });
    window?.runtime?.EventsOn(AppEvents.statusEndKey, () => {
      if (this._task.isCreated) {
        this.remove();
        this.done$.next(true);
      }
    });
  }

  private restoreDefaultTitle() {
    this.linkService.setTitle('Safelock');
  }
}
