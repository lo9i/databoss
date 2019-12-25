import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AgGridModule} from 'ag-grid-angular';
import {ButtonRendererComponent} from './button-renderer/button-renderer.component';
import {ClickableParentComponent} from './clickable-parent/clickable-parent.component';

@NgModule({
  declarations: [ButtonRendererComponent, ClickableParentComponent],
  imports: [
    CommonModule,
    AgGridModule.withComponents([]),
  ],
  exports: [
    ClickableParentComponent,
    ButtonRendererComponent
  ]
})
export class UtilsModule {
}
