import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { of } from 'rxjs';

import { LinkService } from '../../services/link.service';
import { AboutComponent } from './about.component';

describe('AboutComponent', () => {
  let component: AboutComponent;
  let fixture: ComponentFixture<AboutComponent>;

  const version = '12.12.12';

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [AboutComponent],
      providers: [
        { provide: LinkService, useValue: { appVersion$: of(version) }},
      ],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AboutComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });


  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should match static snapshot', () => {
    expect(fixture.nativeElement).toMatchSnapshot();
  });
});
