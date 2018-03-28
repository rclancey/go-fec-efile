from collections import namedtuple

_Name = namedtuple('_Name', ('last', 'first', 'middle', 'prefix', 'suffix'))

_Address = namedtuple('_Address', ('street1', 'street2', 'city', 'state', 'zipcode'))

_Contributor = namedtuple('Contributor', (

class jsonable(object):
    def as_json(self):
        return dict(zip(x, getattr(self, x)) for x in self.__slots__)

    def fields(self):
        data = []
        for x in self.__slots__:
            v = getattr(self, x)
            if isinstance(v, jsonable):
                data.extend(v.fields())
            else:
                data.append(v)
        return data

    @classmethod
    def from_fields(cls, fields):
        data = list(fields[i] for x in xrange(len(cls.__slots__)))
        return cls(*data)

class Name(jsonable):
    __slots__ = ['last', 'first', 'middle', 'prefix', 'suffix']
    def __init__(self, last=None, first=None, middle=None, prefix=None, suffix=None):
        self.last = last
        self.first = first
        self.middle = middle
        self.prefix = prefix
        self.suffix = suffix

class Address(jsonable):
    __slots__ = ['street1', 'street2', 'city', 'state', 'zipcode']
    def __init__(self, street1=None, street2=None, city=None, state=None, zipcode=None):
        self.street1 = street1
        self.street2 = street2
        self.city = city
        self.state = state
        self.zipcode = zipcode

class CounterParty(jsonable):
    __slots__ = ['type', 'organization', 'name', 'address']
    def __init__(self, type=None, organization=None, name=None, address=None):
        self.type = type
        self.organization = organization
        self.name = name
        self.address = address

class Contributor(CounterParty):
    __slots__ = ['type', 'organization', 'name', 'address', 'employer', 'occupation']
    def __init__(self, type=None, organization=None, name=None, address=None, employer=None, occupation=None):
        self.type = type
        self.organization = organization
        self.name = name
        self.address = address
        self.employer = employer
        self.occupaation = occupation

class Committee(jsonable):
    __slots__ = ['fec_id', 'name']
    def __init__(self, fec_id=None, name=None):
        self.fec_id = fec_id
        self.name = name

class Candidate(jsonable):
    __slots__ = ['fec_id', 'name', 'office', 'state', 'district']
    def __init__(self, fec_id=None, name=None, office=None, state=None, district=None):
        self.fec_id = fec_id
        self.name = name
        self.office = office
        self.state = state
        self.district = district

class FECEntity(jsonable):
    __slots__ = ['committee', 'candidate', 'conduit']
    def __init__(self, committee=None, candidate=None, conduit=None):
        self.committee = committee
        self.candidate = candidate
        self.conduit = conduit

            'committee': {
                'fec_id': record[25],
                'name': record[26],
            },
            'candidate': {
                'fec_id': record[27],
                'name': parse_name(record[28:]),
                'office': record[33],
                'state': record[34],
                'district': int_or_none(record[35]),
            },
            'conduit': {
                'name': record[36],
                'address': parse_address(record[37:]),
            },

class Transaction(object):
    pass

class Receipt(Transaction):
    pass

class Disbursement(Transaction):
    pass

