import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';
import {Candidate} from '@models/candidates';
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

  getCandidate(id: string): Observable<Candidate> {
    return this.http.get<Candidate>(`${environment.apiUrl}/candidates/${id}`);
  }
}
