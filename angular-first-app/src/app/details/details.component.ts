import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HousingService } from '../housing.service';
import { HousingLocation } from '../housing-location';

@Component({
  standalone: true,
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent {
  housingLocation: HousingLocation | undefined;

  constructor(route: ActivatedRoute, housingService: HousingService) {
    const housingLocationId = Number(route.snapshot.params['id']);
    this.housingLocation = housingService.getHousingLocationById(housingLocationId);
  }
}
