import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { provideRouter } from '@angular/router';

import { TestHelper } from '../../index.spec';
import { HeaderComponent } from './header.component';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;
  let helper: TestHelper<HeaderComponent>;

  const darkModeBtnSelector = '.btn[data-access="dark-mode-btn"]';
  const aboutBtnSelector = '.btn[data-access="about-btn"]';

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [HeaderComponent],
      providers: [
        provideRouter([
          { path: 'about', component: HeaderComponent },
        ]),
      ],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    helper = new TestHelper(fixture);
    fixture.detectChanges();
  });

  it('should create header component', () => {
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

