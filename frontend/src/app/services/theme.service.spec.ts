import { TestBed } from '@angular/core/testing';

import { ThemeService } from './theme.service';

describe('ThemeService', () => {
  let service: ThemeService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ThemeService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it.each([[false, true], [true, false], [false, true], [true, false]])(
    'should toggle dark mode (expected before: %i) (expected after: %i)',
    (inBefore, inAfter) => {
      expect(service.inDarkMode).toEqual(inBefore);
      service.toggleDarkMode();
      expect(service.inDarkMode).toEqual(inAfter);
    }
  );
});
