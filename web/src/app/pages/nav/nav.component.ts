import {Component, OnInit} from '@angular/core';
import {BreakpointObserver, Breakpoints} from '@angular/cdk/layout';
import {AuthService} from '@services';
import {map} from 'rxjs/operators';
import {Observable} from 'rxjs';
import {Router} from '@angular/router';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.scss']
})
export class NavComponent {
  constructor(private breakpointObserver: BreakpointObserver, private auth: AuthService, private router: Router) {

  }

  isHandset$: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Handset)
    .pipe(
      map(result => result.matches)
    );

  userData: Observable<any> = this.auth.getUserData();

  logout() {
    console.log('loging out');
    this.auth.logout().subscribe(success => {
      this.router.navigate(['/landing']);
    });
  }
}
