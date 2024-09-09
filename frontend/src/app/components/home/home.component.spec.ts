import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { Subject } from 'rxjs';

import * as app from '../../../../wailsjs/go/backend/App';
import { LinkService } from '../../services/link.service';
import { TestHelper } from '../../index.spec';

import { HomeComponent } from './home.component';


describe('HomeComponent', () => {
  let component: HomeComponent;
  let helper: TestHelper<HomeComponent>;
  let fixture: ComponentFixture<HomeComponent>;

  const homeSelector = 'div[data-access="home"]';
  const passwordSelector = 'div[data-access="password"]';
  const processingSelector = 'div[data-access="processing"]';
  const runTaskBtnSelector = 'button[data-access="run-task-btn"]';
  const cancelTaskBtnSelector = 'button[data-access="cancel-task-btn"]';

  const mockLinkService = {
    files$: new Subject(),
    removeDroppedListeners: jest.fn(),
    showError: jest.fn(),
    setTitle: jest.fn(),
  };

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [HomeComponent],
      providers: [
        { provide: LinkService, useValue: mockLinkService },
      ],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    helper = new TestHelper(fixture);
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should show home when files not dropped yet', () => {
    expect(component.taskService.task.isCreated).toBeFalsy();
    expect(helper.select(homeSelector)).toBeTruthy();
  });

  it('should show password input when valid files dropped', () => {
    component.taskService.remove()
    mockLinkService.files$.next(["temp/testing.txt"]);
    fixture.detectChanges();

    const pwd = helper.select(passwordSelector);

    expect(component.taskService.task.isCreated).toBeTruthy();
    expect(component.taskService.task.isRunning).toBeFalsy();
    expect(pwd).toBeTruthy();
  });

  it('should show error when multiple .sla files dropped', () => {
    component.taskService.remove()
    mockLinkService.files$.next(["temp/1.sla", "temp/2.sla"]);
    fixture.detectChanges();

    expect(component.taskService.task.isCreated).toBeFalsy();
    expect(component.taskService.task.isRunning).toBeFalsy();
    expect(mockLinkService.showError).toHaveBeenCalled();
    expect(helper.select(homeSelector)).toBeTruthy();
  });

  it('should prevent passwords shorter than 8 chars from running task', () => {
    component.taskService.remove()
    mockLinkService.files$.next(["temp/testing.txt"]);
    fixture.detectChanges();
    component.password.setValue('short');
    fixture.detectChanges();
    helper.select<HTMLButtonElement>(runTaskBtnSelector)?.click();
    fixture.detectChanges();

    const pwdInput = helper.select(passwordSelector + '> input');

    expect(component.taskService.task.isCreated).toBeTruthy();
    expect(component.taskService.task.isRunning).toBeFalsy();
    expect(pwdInput).toBeTruthy();
    expect(pwdInput?.classList || []).toContain('is-invalid');
  });

  it('should run task when password is valid', () => {
    jest.spyOn(app, "Encrypt").mockReturnValue(new Promise(r => r("test-id")));
    component.taskService.remove()
    mockLinkService.files$.next(["temp/testing.txt"]);
    fixture.detectChanges();
    component.password.setValue('123456789');
    fixture.detectChanges();
    helper.select<HTMLButtonElement>(runTaskBtnSelector)?.click();
    fixture.detectChanges();

    expect(component.taskService.task.isCreated).toBeTruthy();
    expect(component.taskService.task.isRunning).toBeTruthy();
  });

  it('should cancel task and go home when cancel btn clicked', () => {
    component.taskService.remove()
    mockLinkService.files$.next(["temp/testing.txt"]);
    fixture.detectChanges();
    component.password.setValue('123456789');
    fixture.detectChanges();
    helper.select<HTMLButtonElement>(cancelTaskBtnSelector)?.click();
    fixture.detectChanges();

    expect(component.taskService.task.isCreated).toBeFalsy();
    expect(component.taskService.task.isRunning).toBeFalsy();
    expect(helper.select(homeSelector)).toBeTruthy();
  });

  it('should display progress when task is running', () => {
    jest.spyOn(app, "Encrypt").mockReturnValue(new Promise(r => r("test-id")));
    component.taskService.remove()
    mockLinkService.files$.next(["temp/testing.txt"]);
    fixture.detectChanges();
    component.password.setValue('123456789');
    fixture.detectChanges();
    helper.select<HTMLButtonElement>(runTaskBtnSelector)?.click();
    fixture.detectChanges();

    const status = "Encrypting files!"
    const percent = "11.22";
    component.taskService.task.status$.next(status);
    component.taskService.task.percent$.next(percent);
    fixture.detectChanges();
    const processing = helper.select(processingSelector);

    expect(component.taskService.task.isCreated).toBeTruthy();
    expect(component.taskService.task.isRunning).toBeTruthy();
    expect(processing).toBeTruthy();
    expect(processing?.textContent).toContain(status);
    expect(processing?.textContent).toContain(`${percent}%`);
  });

  it('should cancel task when container double clicked while running', () => {
    jest.spyOn(app, "Encrypt").mockReturnValue(new Promise(r => r("test-id")));
    jest.spyOn(app, "Cancel").mockReturnValue(new Promise(r => r()));
    component.taskService.remove()
    mockLinkService.files$.next(["temp/testing.txt"]);
    fixture.detectChanges();
    component.password.setValue('123456789');
    fixture.detectChanges();
    helper.select<HTMLButtonElement>(runTaskBtnSelector)?.click();
    fixture.detectChanges();
    helper.select<HTMLDivElement>('div.container')?.dispatchEvent(new Event("dblclick"));
    fixture.detectChanges();

    expect(component.taskService.task.isCreated).toBeFalsy();
    expect(component.taskService.task.isRunning).toBeFalsy();
    expect(helper.select(homeSelector)).toBeTruthy();
  });
});
