import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CardItemLinkComponent } from './card-item-link.component';

describe('CardItemLinkComponent', () => {
  let component: CardItemLinkComponent;
  let fixture: ComponentFixture<CardItemLinkComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CardItemLinkComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CardItemLinkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
