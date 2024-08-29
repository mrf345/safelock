import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { provideRouter } from '@angular/router';

import { AppComponent } from './app.component';
import { TestHelper } from './index.spec';

describe('AppComponent', () => {
  let component: AppComponent;
  let fixture: ComponentFixture<AppComponent>;
  let helper: TestHelper<AppComponent>;

  const darkModeBtnSelector = '.btn[data-access="dark-mode-btn"]';
  const aboutBtnSelector = '.btn[data-access="about-btn"]';

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [AppComponent],
      providers: [
        provideRouter([
          { path: 'about', component: AppComponent },
        ]),
      ],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
    helper = new TestHelper(fixture);
    fixture.detectChanges();
  });

  it('should create the app', () => {
    expect(component).toBeTruthy();
  });

  it('should toggle dark mode, when dark-mode btn clicked', () => {
    const btn = helper.select<HTMLButtonElement>(darkModeBtnSelector);
    const classesBefore = Array.from(btn?.classList || []);

    btn?.click();
    fixture.detectChanges()
    const classesAfter = Array.from(btn?.classList || []);

    expect(classesBefore.includes('active')).toBeFalsy();
    expect(classesAfter.includes('active')).toBeTruthy();
  });

  it('should activate about btn, when route in about', () => {
    const btn = helper.select<HTMLButtonElement>(aboutBtnSelector);
    const classesBefore = Array.from(btn?.classList || []);

    btn?.click();
    fixture.detectChanges()
    const classesAfter = Array.from(btn?.classList || []);

    expect(classesBefore.includes('active')).toBeFalsy();
    expect(classesAfter.includes('active')).toBeTruthy();
  });
});
