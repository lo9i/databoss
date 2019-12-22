import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CandidatesGridComponent } from './candidates-grid/candidates-grid.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import {
  MatButtonModule,
  MatDatepickerModule,
  MatFormFieldModule,
  MatInputModule,
  MatNativeDateModule,
  MatSelectModule
} from '@angular/material';
import { AgGridModule } from 'ag-grid-angular';


@NgModule({
  declarations: [CandidatesGridComponent],
  imports: [
    CommonModule,
    FormsModule,
    FlexLayoutModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatButtonModule,
    ReactiveFormsModule,
    AgGridModule.withComponents([]),
  ]
})
export class CandidatesModule { }
