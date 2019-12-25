import {Component, OnInit, Input, Output, EventEmitter} from '@angular/core';

@Component({
  selector: 'ag-clickable',
  template: `<button class="{{btnClass}}" mat-raised-button color="accent" type="button" (click)="click($event)">{{label}}</button>`
})

export class ButtonRendererComponent implements OnInit {
  @Input() cell: any;
  @Output() onClicked = new EventEmitter<boolean>();
  params;
  label: string;
  btnClass: string;

  ngOnInit() {
    // this.params = params;
//    this.btnClass = this.params.btnClass || 'btn btn-primary';
    this.label = 'DALE'; // this.params.label || null;
  }

  click(event): void {
    this.onClicked.emit(this.cell);
  }
}
