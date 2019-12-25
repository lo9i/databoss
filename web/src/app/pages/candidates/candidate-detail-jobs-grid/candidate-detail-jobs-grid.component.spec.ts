import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CandidateDetailJobsGridComponent } from './candidate-detail-jobs-grid.component';

describe('CandidateDetailJobsGridComponent', () => {
  let component: CandidateDetailJobsGridComponent;
  let fixture: ComponentFixture<CandidateDetailJobsGridComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CandidateDetailJobsGridComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CandidateDetailJobsGridComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
