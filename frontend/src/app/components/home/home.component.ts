import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { filter, mergeMap, Subscription, zip } from 'rxjs';

import { LinkService } from '../../services/link.service';
import { TaskService } from '../../services/task.service';


@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent implements OnInit, OnDestroy {
  readonly subs = new Subscription();
  readonly password = new FormControl('', {
      validators: [
        Validators.minLength(8),
        Validators.required,
      ],
      updateOn: 'change',
  });

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
