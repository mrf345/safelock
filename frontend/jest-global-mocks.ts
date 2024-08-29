Object.defineProperty(document, 'doctype', {
  value: '<!DOCTYPE html>',
});

Object.defineProperty(window, 'getComputedStyle', {
  value: () => {
    return {
      display: 'none',
      appearance: ['-webkit-appearance'],
    };
  },
});

/**
 * ISSUE: https://github.com/angular/material2/issues/7101
 * Workaround for JSDOM missing transform property
 */
Object.defineProperty(document.body.style, 'transform', {
  value: () => {
    return {
      enumerable: true,
      configurable: true,
    };
  },
});

Object.defineProperty(window, 'go', {
  value: {
    main: {
      App: {
        GetVersion: jest.fn(),
        Encrypt: jest.fn(),
        Decrypt: jest.fn(),
        Cancel: jest.fn(),
      },
    },
  },
});

Object.defineProperty(window, 'runtime', {
  value: {
    EventsOn: jest.fn(),
    EventsOff: jest.fn(),
    OnFileDrop: jest.fn(),
    OnFileDropOff: jest.fn(),
    BrowserOpenURL: jest.fn(),
    WindowSetTitle: jest.fn(),
  },
});
