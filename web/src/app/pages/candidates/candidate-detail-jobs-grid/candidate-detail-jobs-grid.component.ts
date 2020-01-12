import {Component, Input, OnInit} from '@angular/core';
import {GridOptions} from 'ag-grid-community';
import {CandidatesService} from '@services';
import {ActivatedRoute, Router} from '@angular/router';
import * as moment from 'moment';

@Component({
  selector: 'app-candidate-detail-jobs-grid',
  templateUrl: './candidate-detail-jobs-grid.component.html',
  styleUrls: ['./candidate-detail-jobs-grid.component.css']
})
export class CandidateDetailJobsGridComponent implements OnInit {
  public gridOptions: GridOptions;
  public domLayout = 'autoHeight';
  @Input() userId: string;

  constructor(private route: ActivatedRoute, private router: Router, private service: CandidatesService) {
    this.gridOptions = {
      columnDefs: [
        {headerName: 'Id', field: 'name'},
        {headerName: 'Desde', field: 'from', cellRenderer: data => moment(data.from).format('DD/MM/YYYY')},
        {headerName: 'Hasta', field: 'to', cellRenderer: data => moment(data.to).format('DD/MM/YYYY')},
      ],
      defaultColDef: {
        resizable: true,
        filter: true,
        sortable: true,
      },
      rowData: [],
      onGridReady: (params) => this.onGridReady(params)
    };
  }

  onGridReady(params) {
    this.fetchRows();
  }

  ngOnInit() {
  }

  public fetchRows() {
    this.service.getCandidateJobs(this.userId).
    subscribe(data => {
      this.gridOptions.rowData = data;
      setTimeout(() => this.gridOptions.api.sizeColumnsToFit(), 0);
    });
  }
}
