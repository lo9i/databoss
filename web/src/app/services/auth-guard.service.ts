import {Injectable} from '@angular/core';
import {CanActivate, Router} from '@angular/router';
import {AuthService} from './auth.service';

@Injectable()
export class AuthGuardService implements CanActivate {

  constructor(private authService: AuthService, private router: Router) {
  }

  canActivate() {
    if (!this.authService.isLoggedIn()) {
      console.log('canActivate NO');
      this.router.navigate(['/unauthorized']);
    }
    console.log('canActivate SI');
    return this.authService.isLoggedIn();
  }
}
