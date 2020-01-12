import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AgGridModule} from 'ag-grid-angular';
import {ButtonRendererComponent} from './button-renderer/button-renderer.component';
import {ClickableParentComponent} from './clickable-parent/clickable-parent.component';
import {SearchBoxComponent} from './search-box/search-box.component';
import {MatIconModule} from '@angular/material';
import {FormsModule} from '@angular/forms';

@NgModule({
  declarations: [ButtonRendererComponent, ClickableParentComponent, SearchBoxComponent],
  imports: [
    CommonModule,
    MatIconModule,
    FormsModule,
    AgGridModule.withComponents([]),
  ],
  exports: [
    ClickableParentComponent,
    ButtonRendererComponent,
    SearchBoxComponent
  ]
})
export class UtilsModule {
}
