import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TopbarMyAccountComponent } from './topbar-my-account.component';

describe('TopbarMyAccountComponent', () => {
  let component: TopbarMyAccountComponent;
  let fixture: ComponentFixture<TopbarMyAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TopbarMyAccountComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(TopbarMyAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
