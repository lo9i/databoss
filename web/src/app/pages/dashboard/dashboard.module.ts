import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardLayoutComponent } from './dashboard-layout/dashboard-layout.component';
import { FlexLayoutModule } from '@angular/flex-layout';


@NgModule({
  declarations: [DashboardLayoutComponent],
  imports: [
    CommonModule,
    FlexLayoutModule,
  ]
})
export class DashboardModule { }
