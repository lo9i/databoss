import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {CandidatesGridComponent} from './candidates-grid/candidates-grid.component';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FlexLayoutModule} from '@angular/flex-layout';
import {
  MatButtonModule, MatCardModule,
  MatDatepickerModule, MatDividerModule,
  MatFormFieldModule,
  MatInputModule, MatListModule,
  MatNativeDateModule,
  MatSelectModule
} from '@angular/material';
import {AgGridModule} from 'ag-grid-angular';
import {CandidateDetailComponent} from '@pages/candidates/candidate-detail/candidate-detail.component';
import {ClickableParentComponent} from '@utils/clickable-parent/clickable-parent.component';
import {UtilsModule} from '@utils';
import { CandidateDetailJobsGridComponent } from './candidate-detail-jobs-grid/candidate-detail-jobs-grid.component';


@NgModule({
  declarations: [CandidatesGridComponent, CandidateDetailComponent, CandidateDetailJobsGridComponent],
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
    MatDividerModule,
    MatListModule,
    MatCardModule,
    ReactiveFormsModule,
    AgGridModule.withComponents([ClickableParentComponent]),
    UtilsModule
  ]
})
export class CandidatesModule {
}
