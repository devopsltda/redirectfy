import { TestBed } from '@angular/core/testing';

import { RedirectifyApiService } from './redirectify-api.service';

describe('RedirectifyApiService', () => {
  let service: RedirectifyApiService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RedirectifyApiService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
