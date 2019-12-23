import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';
import {Candidate} from '@models/candidates';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CandidatesService {

  constructor(private http: HttpClient) { }
  getCandidates(): Observable<Array<Candidate>> {
    console.log(`${environment.apiUrl}/candidates`);
    return this.http.get<Array<Candidate>>(`${environment.apiUrl}/candidates`);
  }
}
