import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirectifyLogoNameComponent } from './redirectify-logo-name.component';

describe('RedirectifyLogoNameComponent', () => {
  let component: RedirectifyLogoNameComponent;
  let fixture: ComponentFixture<RedirectifyLogoNameComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RedirectifyLogoNameComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RedirectifyLogoNameComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
