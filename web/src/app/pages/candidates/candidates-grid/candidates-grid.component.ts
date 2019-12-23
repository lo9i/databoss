import {Component} from '@angular/core';
import {GridOptions} from 'ag-grid-community';
import {CandidatesService} from '@services';

@Component({
  selector: 'app-candidates-grid',
  templateUrl: './candidates-grid.component.html',
  styleUrls: ['./candidates-grid.component.scss']
})
export class CandidatesGridComponent {
  public gridOptions: GridOptions;

  constructor(private service: CandidatesService) {
    this.gridOptions = {
      columnDefs: [
        {headerName: 'Id', field: 'id', pinned: 'left', maxWidth: 80, resizable: false},
        {headerName: 'Nombre', field: 'firstName'},
        {headerName: 'Apellido', field: 'lastName'},
      ],
      groupDefaultExpanded: -1,
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

  autoSizeAll() {
    this.gridOptions.api.sizeColumnsToFit();
  }

  public fetchRows() {
    this.service.getCandidates().
    subscribe(data => {
      this.gridOptions.rowData = data;
      setTimeout(() => this.autoSizeAll(), 0);
    });
  }
}
