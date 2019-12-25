import {Component} from '@angular/core';
import {ICellRendererAngularComp} from 'ag-grid-angular';

@Component({
  selector: 'app-clickable-parent',
  template: `
        <ag-clickable (onClicked)="clicked($event)" [cell]="cell"></ag-clickable>`
})

export class ClickableParentComponent implements ICellRendererAngularComp {
  private params: any;
  public cell: any;

  agInit(params: any): void {
    this.params = params;
    this.cell = {row: params.value, col: params.colDef.headerName};
  }

  public clicked(cell: any): void {
    this.params.onClicked(cell);
    console.log('Child Cell Clicked: ' + JSON.stringify(cell));
  }

  refresh(): boolean {
    return false;
  }
}
