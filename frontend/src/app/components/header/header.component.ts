import { Component } from '@angular/core';
import { Router, RouterModule, RouterOutlet } from '@angular/router';
import { CommonModule } from '@angular/common';
import { NgIconComponent, provideIcons, provideNgIconsConfig } from '@ng-icons/core';
import { matVisibilityRound, matQuestionMarkRound } from '@ng-icons/material-icons/round';

import { ThemeService } from '../../services/theme.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss',
  standalone: true,
  imports: [
    RouterOutlet,
    CommonModule,
    NgIconComponent,
    RouterModule,
  ],
  providers: [
    provideNgIconsConfig({}),
    provideIcons({
      matVisibilityRound,
      matQuestionMarkRound,
    }),
  ],
})
export class HeaderComponent {
  get inAbout(): boolean {
    return this.router.url === '/about'
  }

  constructor(
    private readonly router: Router,
    readonly themeService: ThemeService,
  ) {}
}
