import { Injectable } from '@angular/core';


@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  readonly darkModeKey = "in-dark-theme";

  get inDarkMode(): boolean {
    return !!localStorage.getItem(this.darkModeKey);
  }

  constructor() {
    this.persistDarkMode()
  }

  private persistDarkMode() {
    if (this.inDarkMode && !document.body.classList.contains(this.darkModeKey)) {
      document.body.classList.add(this.darkModeKey);
    }
  }

  toggleDarkMode() {
    if (this.inDarkMode) {
      localStorage.removeItem(this.darkModeKey);
      document.body.classList.remove(this.darkModeKey);
    } else {
      localStorage.setItem(this.darkModeKey, 'true');
      document.body.classList.add(this.darkModeKey);
    }
  }
}
