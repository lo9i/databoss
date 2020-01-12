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
  hierarchy: string[];
  cardSize = [360, 448];
  cardOrientation = 'horizontal';
  public barChartOptions = {
    borderWidth: 1,
    responsive: true,
    cutoutPercentage: 80,
    backgroundColor: '#666666',
    aspectRatio: 1
  };
  chartsData = [ {color: 'rgba(102, 102, 102, 0.2)', description: 'New Visits', stats: '57,820', icon: 'person'},
    {color: 'rgba(102, 102, 102, 0.2)', description: 'Purchases'   , stats: '$ 89,745', icon: 'money'},
    {color: 'rgba(102, 102, 102, 0.2)', description: 'Active Users', stats: '178,391', icon: 'face'},
    {color: 'rgba(102, 102, 102, 0.2)', description: 'Returned'    , stats: '32,592',  icon: 'refresh'}];


// <!--                  <div class="chart" rel="rgba(102, 102, 102, 0.2)" data-percent="60"><span class="percent">75</span>-->
// <!--                  <i class="chart-icon i-person"></i></div>-->
// <!--        </div>&lt;!&ndash; end ngRepeat: chart in charts &ndash;&gt;-->
//   <!--        <div class="pie-chart-item-container ng-scope" ng-repeat="chart in charts">-->
//   <!--          <div ba-panel="">-->
//   <!--            <div class="panel animated zoomIn" zoom-in="">-->
//   <!--              <div class="panel-body" ng-transclude="">-->
//   <!--                <div class="pie-chart-item ng-scope">-->
// <!--                  <div class="chart" rel="rgba(102, 102, 102, 0.2)" data-percent="60"><span class="percent">79</span>-->
//   <!--                    <canvas height="168" width="168" style="height: 84px; width: 84px;"></canvas>-->
//   <!--                  </div>-->
//   <!--                  <div class="description">-->
//   <!--                    <div class="ng-binding">Purchases</div>-->
//   <!--                    <div class="description-stats ng-binding">$ 89,745</div>-->
// <!--                  </div>-->
// <!--                  <i class="chart-icon i-money"></i></div>-->
// <!--              </div>-->
// <!--            </div>-->
// <!--          </div>-->
// <!--        </div>&lt;!&ndash; end ngRepeat: chart in charts &ndash;&gt;-->
//   <!--        <div class="pie-chart-item-container ng-scope" ng-repeat="chart in charts">-->
//   <!--          <div ba-panel="">-->
//   <!--            <div class="panel animated zoomIn" zoom-in="">-->
//   <!--              <div class="panel-body" ng-transclude="">-->
//   <!--                <div class="pie-chart-item ng-scope">-->
// <!--                  <div class="chart" rel="rgba(102, 102, 102, 0.2)" data-percent="60"><span class="percent">55</span>-->
//   <!--                    <canvas height="168" width="168" style="height: 84px; width: 84px;"></canvas>-->
//   <!--                  </div>-->
//   <!--                  <div class="description">-->
//   <!--                    <div class="ng-binding">Active Users</div>-->
// <!--                    <div class="description-stats ng-binding">178,391</div>-->
// <!--                  </div>-->
// <!--                  <i class="chart-icon i-face"></i></div>-->
// <!--              </div>-->
// <!--            </div>-->
// <!--          </div>-->
// <!--        </div>&lt;!&ndash; end ngRepeat: chart in charts &ndash;&gt;-->
//   <!--        <div class="pie-chart-item-container ng-scope" ng-repeat="chart in charts">-->
//   <!--          <div ba-panel="">-->
//   <!--            <div class="panel animated zoomIn" zoom-in="">-->
//   <!--              <div class="panel-body" ng-transclude="">-->
//   <!--                <div class="pie-chart-item ng-scope">-->
// <!--                  <div class="chart" rel="rgba(102, 102, 102, 0.2)" data-percent="60"><span class="percent">77</span>-->
//   <!--                    <canvas height="168" width="168" style="height: 84px; width: 84px;"></canvas>-->
//   <!--                  </div>-->
//   <!--                  <div class="description">-->
//   <!--                    <div class="ng-binding">Returned</div>-->
//   <!--                    <div class="description-stats ng-binding">32,592</div>-->
// <!--                  </div>-->
// <!--                  <i class="chart-icon i-refresh"></i></div>-->
// <!--              </div>-->
// <!--            </div>-->
// <!--          </div>-->
// <!--        </div>&lt;!&ndash; end ngRepeat: chart in charts &ndash;&gt;</div>-->
// <!--    </dashboard-pie-chart>-->


  constructor(private breakpointObserver: BreakpointObserver,
              private route: ActivatedRoute,
              private ngZone: NgZone,
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
