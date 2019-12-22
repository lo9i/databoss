import * as _ from 'lodash';
import * as moment from 'moment';
import { Component, OnInit, OnDestroy, NgZone } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { ActivatedRoute } from '@angular/router';
import { NgxSpinnerService } from 'ngx-spinner';


@Component({
  selector: 'app-dashboard-layout',
  templateUrl: './dashboard-layout.component.html',
  styleUrls: ['./dashboard-layout.component.scss'],
})
export class DashboardLayoutComponent implements OnInit {
  busy: boolean;
  // timer: NodeJS.Timer;
  hierarchy: string[];
  cardSize = [360, 448];
  cardOrientation = 'horizontal';

  constructor(private breakpointObserver: BreakpointObserver,
              private route: ActivatedRoute,
              private ngZone : NgZone,
              private spinner: NgxSpinnerService) {
    this.makeResponsive();
  }

  makeResponsive() {
    this.breakpointObserver.observe([Breakpoints.HandsetPortrait]).subscribe(result => {
      if (result.matches) {
        this.cardSize = [360, 448];
        this.cardOrientation = 'horizontal';
      }
    });
    this.breakpointObserver.observe([Breakpoints.HandsetLandscape]).subscribe(result => {
      if (result.matches) {
        this.cardSize = [360, 448];
        this.cardOrientation = 'horizontal';
      }
    });
    this.breakpointObserver.observe([Breakpoints.TabletPortrait]).subscribe(result => {
      if (result.matches) {
        this.cardSize = [500, 400];
        this.cardOrientation = 'horizontal';
      }
    });
    this.breakpointObserver.observe([Breakpoints.TabletLandscape]).subscribe(result => {
      if (result.matches) {
        this.cardSize = [800, 400];
        this.cardOrientation = 'vertical';
      }
    });
    this.breakpointObserver.observe([Breakpoints.Web]).subscribe(result => {
      if (result.matches) {
        this.cardSize = [800, 400];
        this.cardOrientation = 'vertical';
      }
    });
  }

  ngOnInit() {
  }
}
