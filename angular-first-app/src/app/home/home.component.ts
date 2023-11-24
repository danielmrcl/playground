import { Component } from '@angular/core';
import { HousingLocationComponent } from '../housing-location/housing-location.component';
import { HousingLocation } from '../housing-location';
import { NgFor } from '@angular/common';
import { HousingService } from '../housing.service';

@Component({
  standalone: true,
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  imports: [
    HousingLocationComponent,
    NgFor
  ]
})
export class HomeComponent {
  housingLocations: HousingLocation[];

  constructor(housingService: HousingService) {
    this.housingLocations = housingService.getAllHousingLocations();
  }
}
