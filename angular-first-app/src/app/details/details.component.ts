import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  standalone: true,
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent {
  housingLocationId: number;

  constructor(route: ActivatedRoute) {
    this.housingLocationId = route.snapshot.params['id'];
  }
}
