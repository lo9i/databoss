import {Routes} from '@angular/router';
import {NavComponent} from './nav/nav.component';
import {LandingComponent} from './landing/landing.component';
import {UnauthorizedComponent} from './unauthorized/unauthorized.component';
// import {DashboardLayoutComponent} from './dashboard/dashboard.component';
// import {DashboardToolbarComponent} from './dashboard/dashboard-toolbar/dashboard-toolbar.component';
import {CandidatesGridComponent} from "./candidates/candidates-grid/candidates-grid.component";
import {AuthGuardService} from "../services/auth-guard.service";

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
  // children: [
  //   {
  // //   path: '',
  // //   component: DashboardToolbarComponent,
  // //   outlet: 'toolbar'
  // // }, {
  //   path: '',
  //   component: DashboardLayoutComponent
  // }]
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
