import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {FlexLayoutModule} from '@angular/flex-layout';
import {Inject, NgModule} from '@angular/core';
import {CookieService} from 'ngx-cookie-service';
import {NgxSpinnerModule} from 'ngx-spinner';
import {HttpClientModule} from '@angular/common/http';

import {AppComponent} from './app.component';
import {PagesModule} from '@pages';
import {ServicesModule} from '@services';
import {UtilsModule} from '@utils';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule.withServerTransition({appId: 'ng-cli-universal'}),
    BrowserAnimationsModule,
    FlexLayoutModule,
    HttpClientModule,
    ServicesModule,
    UtilsModule,
    PagesModule,
    NgxSpinnerModule,
  ],
  providers: [
    CookieService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor() {
  }
}

