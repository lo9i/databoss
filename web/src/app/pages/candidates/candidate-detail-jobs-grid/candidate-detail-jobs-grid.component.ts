import {Component, Input, OnInit} from '@angular/core';
import {GridOptions} from 'ag-grid-community';
import {CandidatesService} from '@services';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-candidate-detail-jobs-grid',
  templateUrl: './candidate-detail-jobs-grid.component.html',
  styleUrls: ['./candidate-detail-jobs-grid.component.css']
})
export class CandidateDetailJobsGridComponent implements OnInit {
  public gridOptions: GridOptions;
  @Input() userId: string;

  constructor(private route: ActivatedRoute, private router: Router, private service: CandidatesService) {
    this.gridOptions = {
      columnDefs: [
        {headerName: 'Nombre', field: 'firstName'},
        {headerName: 'Desde', field: 'lastName'},
        {headerName: 'Hasta', field: 'lastName'},
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
