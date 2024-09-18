import { Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { filter, mergeMap, Subscription, zip } from 'rxjs';

import { LinkService } from '../../services/link.service';
import { TaskService } from '../../services/task.service';
import { getPwdControl, pwdPatternKeys  } from '../../helpers/getPasswordFormControl.function';


@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent implements OnInit, OnDestroy {
  readonly subs = new Subscription();
  readonly pwdMinLength = 8;
  readonly pwdPtrKeys = pwdPatternKeys
  readonly password = getPwdControl(this.pwdMinLength);

  @ViewChild('passwordInput', {}) set passwordInput(input: ElementRef) {
    if (input) setTimeout(() => input.nativeElement.focus(), 1);
  }

  constructor(
    readonly linkService: LinkService,
    readonly taskService: TaskService,
  ) {}

  ngOnInit() {
    this.subs.add(
      zip([
        this.linkService.files$.pipe(filter(fs => !!fs.length)),
        this.taskService.task$.pipe(mergeMap(t => t.notCreated$), filter(Boolean)),
      ]).subscribe(([files]) => this.taskService.create(files))
    );
    this.subs.add(this.taskService.done$
      .pipe(filter(Boolean))
      .subscribe(() => this.cleanup())
    );
  }

  ngOnDestroy() {
    this.subs.unsubscribe()
    this.linkService.removeDroppedListeners();
  }

  runTask() {
    if (this.password.valid && this.taskService.task.isCreated) {
      this.taskService.setPassword(this.password.value as string);
      this.subs.add(this.taskService.run().subscribe(
        id => !id && this.cleanup()
      ));
    }
  }

  cancelTask() {
    if (this.taskService.task.isRunning) {
      this.subs.add(this.taskService.cancel().subscribe());
      this.cleanup();
    } else if (this.taskService.task.isCreated) {
      this.taskService.remove();
      this.cleanup();
    }
  }

  private cleanup() {
    this.password.setValue('');
  }
}
