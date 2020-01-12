import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';
import {CommonModule} from '@angular/common';
import {LayoutModule} from '@angular/cdk/layout';
import {FlexLayoutModule} from '@angular/flex-layout';
import {
  MatToolbarModule,
  MatButtonModule,
  MatSidenavModule,
  MatIconModule,
  MatListModule,
  MatGridListModule,
  MatCardModule,
  MatMenuModule, MatInputModule
} from '@angular/material';
import {routes} from './pages.routes';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {NavComponent} from '@pages/nav/nav.component';
import {DashboardModule} from '@pages/dashboard';
import {CandidatesModule} from '@pages/candidates/candidates.module';
import {LandingComponent} from '@pages/landing/landing.component';
import {UnauthorizedComponent} from '@pages/unauthorized/unauthorized.component';
import {UtilsModule} from '@utils';

@NgModule({
  declarations: [NavComponent, LandingComponent, UnauthorizedComponent],
  imports: [
    RouterModule.forRoot(routes),
    CommonModule,
    LayoutModule,
    FormsModule,
    FlexLayoutModule,
    MatToolbarModule,
    MatButtonModule,
    MatFormFieldModule,
    MatSidenavModule,
    MatIconModule,
    MatListModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatInputModule,
    ReactiveFormsModule,
    UtilsModule,
    DashboardModule,
    CandidatesModule,
  ],
  exports: [RouterModule]
})
export class PagesModule {
}
