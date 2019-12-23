import {Routes} from '@angular/router';
import {NavComponent} from './nav/nav.component';
import {LandingComponent} from './landing/landing.component';
import {UnauthorizedComponent} from './unauthorized/unauthorized.component';
import {CandidatesGridComponent} from './candidates/candidates-grid/candidates-grid.component';
import {AuthGuardService} from '@services';
import {DashboardLayoutComponent} from '@pages/dashboard';

export const routes: Routes = [{
  path: 'landing',
  component: LandingComponent
}, {
  path: 'unauthorized',
  component: UnauthorizedComponent
}, {
  path: 'dashboard',
  component: NavComponent,
  canActivate: [AuthGuardService],
  children: [
    {
      path: '',
      component: DashboardLayoutComponent
    }]
}, {
  path: 'candidates',
  component: NavComponent,
  canActivate: [AuthGuardService],
  children: [{
    path: '',
    component: CandidatesGridComponent
  }]
},
  {
    path: '**',
    redirectTo: '/landing'
  }
];
