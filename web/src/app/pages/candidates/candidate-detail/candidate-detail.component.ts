import {Component, OnInit} from '@angular/core';
import {Candidate} from '@models/candidates';
import {ActivatedRoute} from '@angular/router';
import {CandidatesService} from '@services';
import {Observable} from 'rxjs';

@Component({
  selector: 'app-candidate-detail',
  templateUrl: './candidate-detail.component.html',
  styleUrls: ['./candidate-detail.component.scss']
})
export class CandidateDetailComponent implements OnInit {

  public candidate: Candidate;
  candidateId: string;

  constructor(private route: ActivatedRoute, private service: CandidatesService) {
  }

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      this.candidateId = this.route.snapshot.params.userId;
      this.service.getCandidateByUserId(this.candidateId).
      subscribe(data => {
        this.candidate = data;
      });
    });
  }
}
