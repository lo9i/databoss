import {Component, OnInit} from '@angular/core';
import {GridOptions} from 'ag-grid-community';
import {CandidatesService} from '@services';
import {ActivatedRoute, Router} from '@angular/router';
import {ClickableParentComponent} from '../../../utils/clickable-parent/clickable-parent.component';

@Component({
  selector: 'app-candidates-grid',
  templateUrl: './candidates-grid.component.html',
  styleUrls: ['./candidates-grid.component.scss']
})
export class CandidatesGridComponent implements OnInit {
  public gridOptions: GridOptions;
  searchText: string;

  constructor(private route: ActivatedRoute, private router: Router, private service: CandidatesService) {
    this.gridOptions = {
      columnDefs: [
        {headerName: 'CUIL/T', field: 'userId', pinned: 'left', maxWidth: 180, resizable: false,  cellRendererFramework: ClickableParentComponent,
          cellRendererParams: {
            onClicked: this.detail.bind(this),
            btnClass: 'primary',
            label: 'Detalle'
          }},
        {headerName: 'Nombre', field: 'firstName'},
        {headerName: 'Apellido', field: 'lastName'},
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
    this.route.paramMap.subscribe(params => {
      this.searchText = this.route.snapshot.queryParamMap.get('userId');
    });
  }
  public fetchRows() {
    this.service.getCandidates(this.searchText).
    subscribe(data => {
      this.gridOptions.rowData = data;
      setTimeout(() => this.gridOptions.api.sizeColumnsToFit(), 0);
    });
  }

  detail(cell) {
    this.router.navigate(['/candidate-detail/' + cell.row ]);
  }
}
