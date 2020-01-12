import {Component, Input} from '@angular/core';
import {Candidate} from '@models/candidates';

@Component({
  selector: 'app-candidate-detail-header',
  templateUrl: './candidate-detail-header.component.html',
  styleUrls: ['./candidate-detail-header.component.css']
})
export class CandidateDetailHeaderComponent {
  @Input() candidate: Candidate;
  constructor() { }
}
