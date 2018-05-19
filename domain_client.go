package lingo

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// A DomainClient implements the Domainer interface and provides all of the
// functionality for managing domains and DNS records in a Linode account.
type DomainClient struct {
	api APIClient
}

// NewDomainClient returns a new DomainClient given and APIClient.
func NewDomainClient(api APIClient) DomainClient {
	return DomainClient{api: api}
}

// ListDomains retrieves a slice of Domains available to a Linode account.
func (c DomainClient) ListDomains() ([]Domain, error) {
	data, err := c.api.Get("domains")
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListDomains")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListDomains response")
	}

	// TODO: Do something with paging here?
	var domains []Domain
	if err := json.Unmarshal(results.Data, &domains); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListDomains data")
	}

	return domains, nil
}

// ViewDomain retrieves a specific Linode Domain.
func (c DomainClient) ViewDomain(id uint) (Domain, error) {
	var domain Domain

	data, err := c.api.Get(fmt.Sprintf("domains/%d", id))
	if err != nil {
		return domain, errors.Wrap(err, "failed to make request for ViewDomain")
	}

	if err := json.Unmarshal(data, &domain); err != nil {
		return domain, errors.Wrap(err, "failed to unmarshal ViewDomain data")
	}

	return domain, nil
}

// CreateDomain creates a new Domain in a Linode account.
func (c DomainClient) CreateDomain(domain Domain) (Domain, error) {
	var created Domain

	payload, err := json.Marshal(&domain)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateDomain")
	}

	data, err := c.api.Post("domains", payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for CreateDomain")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to decode CreateDomain response")
	}

	return created, nil
}

// UpdateDomain updates a specific Domain in a Linode account if it exists.
func (c DomainClient) UpdateDomain(domain Domain) (Domain, error) {
	var updated Domain

	payload, err := json.Marshal(&domain)
	if err != nil {
		return updated, errors.Wrap(err, "failed to marshal request for UpdateDomain")
	}

	data, err := c.api.Put("domains", payload)
	if err != nil {
		return updated, errors.Wrap(err, "failed to make request for UpdateDomain")
	}

	if err := json.Unmarshal(data, &updated); err != nil {
		return updated, errors.Wrap(err, "failed to decode UpdateDomain response")
	}

	return updated, nil

}

// DeleteDomain deletes a specific Domain from a Linode account.
func (c DomainClient) DeleteDomain(id uint) (Domain, error) {
	var domain Domain

	_, err := c.api.Delete(fmt.Sprintf("domains/%d", id))
	if err != nil {
		return domain, errors.Wrap(err, "failed to make request for DeleteDomain")
	}

	return domain, nil
}

// ListDomainRecords retrieves a slice of Domain Records available within the specified Domain.
func (c DomainClient) ListDomainRecords(domainID uint) ([]DomainRecord, error) {
	data, err := c.api.Get(fmt.Sprintf("domains/%d/records", domainID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request for ListDomainRecords")
	}

	var results Results
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode ListDomainRecords response")
	}

	// TODO: Do something with paging here?
	var records []DomainRecord
	if err := json.Unmarshal(results.Data, &records); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ListDomainRecords data")
	}

	return records, nil
}

// ViewDomainRecord retrieves a specific Linode Domain Record.
func (c DomainClient) ViewDomainRecord(domainID, recordID uint) (DomainRecord, error) {
	var record DomainRecord

	data, err := c.api.Get(fmt.Sprintf("domains/%d/records/%d", domainID, recordID))
	if err != nil {
		return record, errors.Wrap(err, "failed to make request for ViewDomainRecord")
	}

	if err := json.Unmarshal(data, &record); err != nil {
		return record, errors.Wrap(err, "failed to unmarshal ViewDomainRecord data")
	}

	return record, nil
}

// CreateDomainRecord creates a new Domain Record for a given Domain.
func (c DomainClient) CreateDomainRecord(domainID uint, record DomainRecord) (DomainRecord, error) {
	var created DomainRecord

	payload, err := json.Marshal(&record)
	if err != nil {
		return created, errors.Wrap(err, "failed to marshal request for CreateDomainRecord")
	}

	data, err := c.api.Post(fmt.Sprintf("domains/%d/records", domainID), payload)
	if err != nil {
		return created, errors.Wrap(err, "failed to make request for CreateDomainRecord")
	}

	if err := json.Unmarshal(data, &created); err != nil {
		return created, errors.Wrap(err, "failed to decode CreateDomainRecord response")
	}

	return created, nil
}

// UpdateDomainRecord updates a specific Domain Record in the specified Domain if it exists.
func (c DomainClient) UpdateDomainRecord(domainID uint, record DomainRecord) (DomainRecord, error) {
	var updated DomainRecord

	payload, err := json.Marshal(&record)
	if err != nil {
		return updated, errors.Wrap(err, "failed to marshal request for UpdateDomainRecord")
	}

	data, err := c.api.Put(fmt.Sprintf("domains/%d/records", domainID), payload)
	if err != nil {
		return updated, errors.Wrap(err, "failed to make request for UpdateDomainRecord")
	}

	if err := json.Unmarshal(data, &updated); err != nil {
		return updated, errors.Wrap(err, "failed to decode UpdateDomainRecord response")
	}

	return updated, nil
}

// DeleteDomainRecord deletes a specific Domain Record from the specified Domain.
func (c DomainClient) DeleteDomainRecord(domainID, recordID uint) error {
	if _, err := c.api.Delete(fmt.Sprintf("domains/%d/records/%d", domainID, recordID)); err != nil {
		return errors.Wrap(err, "failed to make request for DeleteDomainRecord")
	}

	return nil
}
