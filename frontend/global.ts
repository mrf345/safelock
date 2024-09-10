export {}

declare global {
  interface Window {
    runtime: {
      EventsOn(eventName: string, callback: (...optionalData: any[]) => void): void,
      EventsOff(eventName: string, ...additionalEvents: string[]): void
      OnFileDrop(callback: (x: number, y: number, paths: string[]) => void, useDropTarget: boolean) :void
      OnFileDropOff(): void
      BrowserOpenURL(url: string): void
    }
  }
}
