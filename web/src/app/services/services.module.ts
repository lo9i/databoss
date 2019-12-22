import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthInterceptor } from './auth.interceptor';
import { AuthService } from './auth.service';
import { AuthGuardService } from './auth-guard.service';
import { CandidatesService } from './candidates.service';


const HttpAuthInterceptor = {
  provide: HTTP_INTERCEPTORS,
  useClass: AuthInterceptor,
  multi: true
};

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    HttpClientModule
  ],
  providers: [
    HttpAuthInterceptor,
    AuthService,
    AuthGuardService,
    CandidatesService,
  ]
})
export class ServicesModule { }
