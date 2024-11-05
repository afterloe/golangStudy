// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: user_service.proto

#include "user_service.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
#include "google/protobuf/generated_message_tctable_impl.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::google::protobuf;
namespace _pbi = ::google::protobuf::internal;
namespace _fl = ::google::protobuf::internal::field_layout;
namespace Logic {

inline constexpr LoginResponse::Impl_::Impl_(
    ::_pbi::ConstantInitialized) noexcept
      : sessionid_(
            &::google::protobuf::internal::fixed_address_empty_string,
            ::_pbi::ConstantInitialized()),
        accesstoken_(
            &::google::protobuf::internal::fixed_address_empty_string,
            ::_pbi::ConstantInitialized()),
        refreshtoken_(
            &::google::protobuf::internal::fixed_address_empty_string,
            ::_pbi::ConstantInitialized()),
        _cached_size_{0} {}

template <typename>
PROTOBUF_CONSTEXPR LoginResponse::LoginResponse(::_pbi::ConstantInitialized)
    : _impl_(::_pbi::ConstantInitialized()) {}
struct LoginResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR LoginResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~LoginResponseDefaultTypeInternal() {}
  union {
    LoginResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 LoginResponseDefaultTypeInternal _LoginResponse_default_instance_;

inline constexpr LoginRequest::Impl_::Impl_(
    ::_pbi::ConstantInitialized) noexcept
      : loginname_(
            &::google::protobuf::internal::fixed_address_empty_string,
            ::_pbi::ConstantInitialized()),
        scrip_(
            &::google::protobuf::internal::fixed_address_empty_string,
            ::_pbi::ConstantInitialized()),
        _cached_size_{0} {}

template <typename>
PROTOBUF_CONSTEXPR LoginRequest::LoginRequest(::_pbi::ConstantInitialized)
    : _impl_(::_pbi::ConstantInitialized()) {}
struct LoginRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR LoginRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~LoginRequestDefaultTypeInternal() {}
  union {
    LoginRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 LoginRequestDefaultTypeInternal _LoginRequest_default_instance_;
}  // namespace Logic
static ::_pb::Metadata file_level_metadata_user_5fservice_2eproto[2];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_user_5fservice_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_user_5fservice_2eproto = nullptr;
const ::uint32_t TableStruct_user_5fservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::Logic::LoginRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::Logic::LoginRequest, _impl_.loginname_),
    PROTOBUF_FIELD_OFFSET(::Logic::LoginRequest, _impl_.scrip_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::Logic::LoginResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::Logic::LoginResponse, _impl_.sessionid_),
    PROTOBUF_FIELD_OFFSET(::Logic::LoginResponse, _impl_.accesstoken_),
    PROTOBUF_FIELD_OFFSET(::Logic::LoginResponse, _impl_.refreshtoken_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, -1, -1, sizeof(::Logic::LoginRequest)},
        {10, -1, -1, sizeof(::Logic::LoginResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::Logic::_LoginRequest_default_instance_._instance,
    &::Logic::_LoginResponse_default_instance_._instance,
};
const char descriptor_table_protodef_user_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n\022user_service.proto\022\005Logic\"0\n\014LoginRequ"
    "est\022\021\n\tLoginName\030\001 \001(\t\022\r\n\005Scrip\030\002 \001(\t\"M\n"
    "\rLoginResponse\022\021\n\tSessionID\030\001 \001(\t\022\023\n\013Acc"
    "essToken\030\002 \001(\t\022\024\n\014RefreshToken\030\003 \001(\t2@\n\n"
    "UserServer\0222\n\005Login\022\023.Logic.LoginRequest"
    "\032\024.Logic.LoginResponseB\007Z\005./;pbb\006proto3"
};
static ::absl::once_flag descriptor_table_user_5fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_user_5fservice_2eproto = {
    false,
    false,
    239,
    descriptor_table_protodef_user_5fservice_2eproto,
    "user_service.proto",
    &descriptor_table_user_5fservice_2eproto_once,
    nullptr,
    0,
    2,
    schemas,
    file_default_instances,
    TableStruct_user_5fservice_2eproto::offsets,
    file_level_metadata_user_5fservice_2eproto,
    file_level_enum_descriptors_user_5fservice_2eproto,
    file_level_service_descriptors_user_5fservice_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_user_5fservice_2eproto_getter() {
  return &descriptor_table_user_5fservice_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_user_5fservice_2eproto(&descriptor_table_user_5fservice_2eproto);
namespace Logic {
// ===================================================================

class LoginRequest::_Internal {
 public:
};

LoginRequest::LoginRequest(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:Logic.LoginRequest)
}
inline PROTOBUF_NDEBUG_INLINE LoginRequest::Impl_::Impl_(
    ::google::protobuf::internal::InternalVisibility visibility, ::google::protobuf::Arena* arena,
    const Impl_& from)
      : loginname_(arena, from.loginname_),
        scrip_(arena, from.scrip_),
        _cached_size_{0} {}

LoginRequest::LoginRequest(
    ::google::protobuf::Arena* arena,
    const LoginRequest& from)
    : ::google::protobuf::Message(arena) {
  LoginRequest* const _this = this;
  (void)_this;
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  new (&_impl_) Impl_(internal_visibility(), arena, from._impl_);

  // @@protoc_insertion_point(copy_constructor:Logic.LoginRequest)
}
inline PROTOBUF_NDEBUG_INLINE LoginRequest::Impl_::Impl_(
    ::google::protobuf::internal::InternalVisibility visibility,
    ::google::protobuf::Arena* arena)
      : loginname_(arena),
        scrip_(arena),
        _cached_size_{0} {}

inline void LoginRequest::SharedCtor(::_pb::Arena* arena) {
  new (&_impl_) Impl_(internal_visibility(), arena);
}
LoginRequest::~LoginRequest() {
  // @@protoc_insertion_point(destructor:Logic.LoginRequest)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void LoginRequest::SharedDtor() {
  ABSL_DCHECK(GetArena() == nullptr);
  _impl_.loginname_.Destroy();
  _impl_.scrip_.Destroy();
  _impl_.~Impl_();
}

PROTOBUF_NOINLINE void LoginRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:Logic.LoginRequest)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.loginname_.ClearToEmpty();
  _impl_.scrip_.ClearToEmpty();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* LoginRequest::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<1, 2, 0, 41, 2> LoginRequest::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    2, 8,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967292,  // skipmap
    offsetof(decltype(_table_), field_entries),
    2,  // num_field_entries
    0,  // num_aux_entries
    offsetof(decltype(_table_), field_names),  // no aux_entries
    &_LoginRequest_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // string Scrip = 2;
    {::_pbi::TcParser::FastUS1,
     {18, 63, 0, PROTOBUF_FIELD_OFFSET(LoginRequest, _impl_.scrip_)}},
    // string LoginName = 1;
    {::_pbi::TcParser::FastUS1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(LoginRequest, _impl_.loginname_)}},
  }}, {{
    65535, 65535
  }}, {{
    // string LoginName = 1;
    {PROTOBUF_FIELD_OFFSET(LoginRequest, _impl_.loginname_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
    // string Scrip = 2;
    {PROTOBUF_FIELD_OFFSET(LoginRequest, _impl_.scrip_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
  }},
  // no aux_entries
  {{
    "\22\11\5\0\0\0\0\0"
    "Logic.LoginRequest"
    "LoginName"
    "Scrip"
  }},
};

::uint8_t* LoginRequest::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:Logic.LoginRequest)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // string LoginName = 1;
  if (!this->_internal_loginname().empty()) {
    const std::string& _s = this->_internal_loginname();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "Logic.LoginRequest.LoginName");
    target = stream->WriteStringMaybeAliased(1, _s, target);
  }

  // string Scrip = 2;
  if (!this->_internal_scrip().empty()) {
    const std::string& _s = this->_internal_scrip();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "Logic.LoginRequest.Scrip");
    target = stream->WriteStringMaybeAliased(2, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:Logic.LoginRequest)
  return target;
}

::size_t LoginRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:Logic.LoginRequest)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string LoginName = 1;
  if (!this->_internal_loginname().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_loginname());
  }

  // string Scrip = 2;
  if (!this->_internal_scrip().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_scrip());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData LoginRequest::_class_data_ = {
    LoginRequest::MergeImpl,
    nullptr,  // OnDemandRegisterArenaDtor
};
const ::google::protobuf::Message::ClassData* LoginRequest::GetClassData() const {
  return &_class_data_;
}

void LoginRequest::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<LoginRequest*>(&to_msg);
  auto& from = static_cast<const LoginRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:Logic.LoginRequest)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_loginname().empty()) {
    _this->_internal_set_loginname(from._internal_loginname());
  }
  if (!from._internal_scrip().empty()) {
    _this->_internal_set_scrip(from._internal_scrip());
  }
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void LoginRequest::CopyFrom(const LoginRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:Logic.LoginRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool LoginRequest::IsInitialized() const {
  return true;
}

::_pbi::CachedSize* LoginRequest::AccessCachedSize() const {
  return &_impl_._cached_size_;
}
void LoginRequest::InternalSwap(LoginRequest* PROTOBUF_RESTRICT other) {
  using std::swap;
  auto* arena = GetArena();
  ABSL_DCHECK_EQ(arena, other->GetArena());
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.loginname_, &other->_impl_.loginname_, arena);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.scrip_, &other->_impl_.scrip_, arena);
}

::google::protobuf::Metadata LoginRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_user_5fservice_2eproto_getter, &descriptor_table_user_5fservice_2eproto_once,
      file_level_metadata_user_5fservice_2eproto[0]);
}
// ===================================================================

class LoginResponse::_Internal {
 public:
};

LoginResponse::LoginResponse(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:Logic.LoginResponse)
}
inline PROTOBUF_NDEBUG_INLINE LoginResponse::Impl_::Impl_(
    ::google::protobuf::internal::InternalVisibility visibility, ::google::protobuf::Arena* arena,
    const Impl_& from)
      : sessionid_(arena, from.sessionid_),
        accesstoken_(arena, from.accesstoken_),
        refreshtoken_(arena, from.refreshtoken_),
        _cached_size_{0} {}

LoginResponse::LoginResponse(
    ::google::protobuf::Arena* arena,
    const LoginResponse& from)
    : ::google::protobuf::Message(arena) {
  LoginResponse* const _this = this;
  (void)_this;
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  new (&_impl_) Impl_(internal_visibility(), arena, from._impl_);

  // @@protoc_insertion_point(copy_constructor:Logic.LoginResponse)
}
inline PROTOBUF_NDEBUG_INLINE LoginResponse::Impl_::Impl_(
    ::google::protobuf::internal::InternalVisibility visibility,
    ::google::protobuf::Arena* arena)
      : sessionid_(arena),
        accesstoken_(arena),
        refreshtoken_(arena),
        _cached_size_{0} {}

inline void LoginResponse::SharedCtor(::_pb::Arena* arena) {
  new (&_impl_) Impl_(internal_visibility(), arena);
}
LoginResponse::~LoginResponse() {
  // @@protoc_insertion_point(destructor:Logic.LoginResponse)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void LoginResponse::SharedDtor() {
  ABSL_DCHECK(GetArena() == nullptr);
  _impl_.sessionid_.Destroy();
  _impl_.accesstoken_.Destroy();
  _impl_.refreshtoken_.Destroy();
  _impl_.~Impl_();
}

PROTOBUF_NOINLINE void LoginResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:Logic.LoginResponse)
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.sessionid_.ClearToEmpty();
  _impl_.accesstoken_.ClearToEmpty();
  _impl_.refreshtoken_.ClearToEmpty();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* LoginResponse::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<2, 3, 0, 60, 2> LoginResponse::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    3, 24,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967288,  // skipmap
    offsetof(decltype(_table_), field_entries),
    3,  // num_field_entries
    0,  // num_aux_entries
    offsetof(decltype(_table_), field_names),  // no aux_entries
    &_LoginResponse_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    {::_pbi::TcParser::MiniParse, {}},
    // string SessionID = 1;
    {::_pbi::TcParser::FastUS1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(LoginResponse, _impl_.sessionid_)}},
    // string AccessToken = 2;
    {::_pbi::TcParser::FastUS1,
     {18, 63, 0, PROTOBUF_FIELD_OFFSET(LoginResponse, _impl_.accesstoken_)}},
    // string RefreshToken = 3;
    {::_pbi::TcParser::FastUS1,
     {26, 63, 0, PROTOBUF_FIELD_OFFSET(LoginResponse, _impl_.refreshtoken_)}},
  }}, {{
    65535, 65535
  }}, {{
    // string SessionID = 1;
    {PROTOBUF_FIELD_OFFSET(LoginResponse, _impl_.sessionid_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
    // string AccessToken = 2;
    {PROTOBUF_FIELD_OFFSET(LoginResponse, _impl_.accesstoken_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
    // string RefreshToken = 3;
    {PROTOBUF_FIELD_OFFSET(LoginResponse, _impl_.refreshtoken_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
  }},
  // no aux_entries
  {{
    "\23\11\13\14\0\0\0\0"
    "Logic.LoginResponse"
    "SessionID"
    "AccessToken"
    "RefreshToken"
  }},
};

::uint8_t* LoginResponse::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:Logic.LoginResponse)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // string SessionID = 1;
  if (!this->_internal_sessionid().empty()) {
    const std::string& _s = this->_internal_sessionid();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "Logic.LoginResponse.SessionID");
    target = stream->WriteStringMaybeAliased(1, _s, target);
  }

  // string AccessToken = 2;
  if (!this->_internal_accesstoken().empty()) {
    const std::string& _s = this->_internal_accesstoken();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "Logic.LoginResponse.AccessToken");
    target = stream->WriteStringMaybeAliased(2, _s, target);
  }

  // string RefreshToken = 3;
  if (!this->_internal_refreshtoken().empty()) {
    const std::string& _s = this->_internal_refreshtoken();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "Logic.LoginResponse.RefreshToken");
    target = stream->WriteStringMaybeAliased(3, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:Logic.LoginResponse)
  return target;
}

::size_t LoginResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:Logic.LoginResponse)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string SessionID = 1;
  if (!this->_internal_sessionid().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_sessionid());
  }

  // string AccessToken = 2;
  if (!this->_internal_accesstoken().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_accesstoken());
  }

  // string RefreshToken = 3;
  if (!this->_internal_refreshtoken().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_refreshtoken());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData LoginResponse::_class_data_ = {
    LoginResponse::MergeImpl,
    nullptr,  // OnDemandRegisterArenaDtor
};
const ::google::protobuf::Message::ClassData* LoginResponse::GetClassData() const {
  return &_class_data_;
}

void LoginResponse::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<LoginResponse*>(&to_msg);
  auto& from = static_cast<const LoginResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:Logic.LoginResponse)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_sessionid().empty()) {
    _this->_internal_set_sessionid(from._internal_sessionid());
  }
  if (!from._internal_accesstoken().empty()) {
    _this->_internal_set_accesstoken(from._internal_accesstoken());
  }
  if (!from._internal_refreshtoken().empty()) {
    _this->_internal_set_refreshtoken(from._internal_refreshtoken());
  }
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void LoginResponse::CopyFrom(const LoginResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:Logic.LoginResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool LoginResponse::IsInitialized() const {
  return true;
}

::_pbi::CachedSize* LoginResponse::AccessCachedSize() const {
  return &_impl_._cached_size_;
}
void LoginResponse::InternalSwap(LoginResponse* PROTOBUF_RESTRICT other) {
  using std::swap;
  auto* arena = GetArena();
  ABSL_DCHECK_EQ(arena, other->GetArena());
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.sessionid_, &other->_impl_.sessionid_, arena);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.accesstoken_, &other->_impl_.accesstoken_, arena);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.refreshtoken_, &other->_impl_.refreshtoken_, arena);
}

::google::protobuf::Metadata LoginResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_user_5fservice_2eproto_getter, &descriptor_table_user_5fservice_2eproto_once,
      file_level_metadata_user_5fservice_2eproto[1]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace Logic
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"