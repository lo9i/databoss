import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';
import {Candidate, Job} from '@models/candidates';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CandidatesService {

  constructor(private http: HttpClient) {
  }

  getCandidates(searchText: string): Observable<Array<Candidate>> {
    let url = `${environment.apiUrl}/candidates`;
    if (searchText != null) {
      url += `?userId=${searchText}`;
    }
    return this.http.get<Array<Candidate>>(url);
  }

  getCandidateById(id: string): Observable<Candidate> {
    return this.http.get<Candidate>(`${environment.apiUrl}/candidates/${id}/detail`);
  }

  getCandidateByUserId(id: string): Observable<Candidate> {
    return this.http.get<Candidate>(`${environment.apiUrl}/candidates/${id}`);
  }

  getCandidateJobs(id: string): Observable<Array<Job>> {
    return this.http.get<Array<Job>>(`${environment.apiUrl}/candidates/${id}/jobs`);
  }
}
