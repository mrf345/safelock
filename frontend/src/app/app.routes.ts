import { Routes } from '@angular/router';

export const routes: Routes = [
    { path: '', redirectTo: '/home', pathMatch: 'full' },
    {
        path: 'about',
        loadComponent: () => import('./components/about/about.component')
            .then(m => m.AboutComponent)
    },
    {
        path: 'home',
        loadComponent: () => import('./components/home/home.component')
            .then(m => m.HomeComponent)
    },
];
