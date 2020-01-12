import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {DashboardLayoutComponent} from './dashboard-layout/dashboard-layout.component';
import {FlexLayoutModule} from '@angular/flex-layout';
import {ChartsModule} from 'ng2-charts';
import {MatCardModule, MatIconModule} from '@angular/material';


@NgModule({
  declarations: [DashboardLayoutComponent],
  imports: [
    CommonModule,
    FlexLayoutModule,
    ChartsModule,
    MatCardModule,
    MatIconModule,
  ]
})
export class DashboardModule {
}
