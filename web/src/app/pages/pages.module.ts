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
  MatMenuModule
} from '@angular/material';
import {routes} from './pages.routes';
import {NavComponent} from './nav/nav.component';
import {LandingComponent} from './landing/landing.component';
import {UnauthorizedComponent} from './unauthorized/unauthorized.component';
import {CandidatesModule} from "@pages/candidates/candidates.module";
import {DashboardModule} from "@pages/dashboard";
import {MatFormFieldModule} from "@angular/material/form-field";
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [NavComponent, LandingComponent, UnauthorizedComponent],
  imports: [
    CommonModule,
    LayoutModule,
    FlexLayoutModule,
    MatToolbarModule,
    MatButtonModule,
    MatFormFieldModule,
    MatSidenavModule,
    MatIconModule,
    MatListModule,
    RouterModule.forRoot(routes),
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    ReactiveFormsModule,
    DashboardModule,
    CandidatesModule
  ],
  exports: [RouterModule]
})
export class PagesModule {
}
