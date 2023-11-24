import { Component } from '@angular/core';
import { HousingLocationComponent } from '../housing-location/housing-location.component';
import { HousingLocation } from '../housing-location';
import { CommonModule } from '@angular/common';
import { HousingService } from '../housing.service';

@Component({
  standalone: true,
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  imports: [
    HousingLocationComponent,
    CommonModule
  ]
})
export class HomeComponent {
  housingLocations: HousingLocation[] = [];
  filteredLocations: HousingLocation[] = [];

  constructor(housingService: HousingService) {
    housingService.getAllHousingLocations().then(locations => {
      this.housingLocations = locations;
      this.filteredLocations = locations;
    });
  }

  filterSearch(searchValue: string) {
    const formatString = (str: string) => str.toLowerCase().replaceAll(' ', '');
    this.filteredLocations = this.housingLocations
      .filter(location => formatString(location.city).includes(formatString(searchValue)));
  }
}
