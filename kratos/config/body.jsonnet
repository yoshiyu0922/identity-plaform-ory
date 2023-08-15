function(ctx) {
  userId: ctx.identity.id,
  customData: ctx.transient_payload.custom_data,
  traits: {
    email: ctx.identity.traits.email,
    name: ctx.identity.traits.name,
    newsletterConsent: ctx.identity.traits.consent.newsletter,
  },
}
