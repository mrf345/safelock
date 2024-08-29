import { ComponentFixture } from "@angular/core/testing";

export class TestHelper<T> {
  fixture: ComponentFixture<T>;

  get dom() {
    return this.fixture.nativeElement as HTMLElement;
  }

  constructor(fixture: ComponentFixture<T>) {
    this.fixture = fixture;
  }

  select<E extends Element>(selector: string) {
    return this.dom.querySelector<E>(selector);
  }

  selectAll<E extends Element>(selector: string) {
    return this.dom.querySelectorAll<E>(selector);
  }
}
