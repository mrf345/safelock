import { BehaviorSubject, combineLatest, map } from "rxjs";

export enum TaskKind {
  encrypt = 'encrypt',
  decrypt = 'decrypt',
}
  
export class Task {
  id$ = new BehaviorSubject('');
  kind$ = new BehaviorSubject<`${TaskKind}` | undefined>(undefined);
  files$ = new BehaviorSubject<string[]>([]);
  status$ = new BehaviorSubject('');
  percent$ = new BehaviorSubject('');
  password$ = new BehaviorSubject('');

  isRunning$ = this.id$.pipe(map(Boolean));
  notRunning$ = this.isRunning$.pipe(map(r => !r));
  notComplete$ = this.percent$.pipe(map(p => 100 > parseFloat(p)));
  isCreated$ = combineLatest([this.kind$, this.files$])
    .pipe(map(([a, b]) => !!(a && b.length)));
  notCreated$ = this.isCreated$.pipe(map(c => !c));
  inProgress$ = combineLatest([this.isRunning$, this.notComplete$])
    .pipe(map(([a, b]) => a && b));
  noProgress$ = this.inProgress$.pipe(map(p => !p));

  get isCreated(): boolean {
    return !!(this.kind$.getValue() || this.files$.getValue().length)
  }

  get isRunning(): boolean {
    return !!this.id$.getValue();
  }
}
