import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'app-search-box',
  templateUrl: './search-box.component.html',
  styleUrls: ['./search-box.component.scss']
})
export class SearchBoxComponent implements OnInit {
  searchText: string;

  constructor(private router: Router) { }

  ngOnInit() {
  }
  onKeydown(event) {
    if (event.key === 'Enter') {
      this.startSearch();
    }
  }

  startSearch() {
    if (this.router.url.startsWith('/candidate-detail')) {
      this.router.navigateByUrl('/', {skipLocationChange: true})
        .then(() => this.router.navigate(['/candidate-detail/' + this.searchText]));
    } else {
      this.router.navigate(['/candidate-detail/' + this.searchText]);
    }
  }
}
